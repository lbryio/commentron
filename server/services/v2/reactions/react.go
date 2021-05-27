package reactions

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/lbryio/commentron/helper"

	"github.com/lbryio/commentron/flags"

	"github.com/lbryio/commentron/commentapi"
	"github.com/lbryio/commentron/db"
	"github.com/lbryio/commentron/model"
	"github.com/lbryio/commentron/server/lbry"

	"github.com/lbryio/errors.go"
	"github.com/lbryio/lbry.go/extras/api"
	"github.com/lbryio/lbry.go/v2/extras/util"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// React creates/updates a reaction to a comment
func react(r *http.Request, args *commentapi.ReactArgs, reply *commentapi.ReactResponse) error {

	comments, err := model.Comments(qm.WhereIn(model.CommentColumns.CommentID+" IN ?", util.StringSplitArg(args.CommentIDs, ",")...)).AllG()
	if err != nil {
		return errors.Err(err)
	}
	if len(comments) == 0 {
		return errors.Err("could not find comments(s)")
	}
	var commentIDs []interface{}
	for _, p := range comments {
		commentIDs = append(commentIDs, p.CommentID)
	}
	if len(commentIDs) > 1 {
		return api.StatusError{Err: errors.Err("only one comment id can be passed currently"), Status: http.StatusBadRequest}
	}
	channel, err := model.Channels(model.ChannelWhere.ClaimID.EQ(args.ChannelID)).OneG()
	if errors.Is(err, sql.ErrNoRows) {
		channel = &model.Channel{
			ClaimID: args.ChannelID,
			Name:    args.ChannelName,
		}
		err = nil
		err := channel.InsertG(boil.Infer())
		if err != nil {
			return errors.Err(err)
		}
	}
	err = lbry.ValidateSignature(args.ChannelID, args.Signature, args.SigningTS, args.ChannelName)
	if err != nil {
		return errors.Prefix("could not authenticate channel signature:", err)
	}
	if len(comments) > 1 {
		logrus.Warningf("%d comments reacted to in the same call from ip[%s] for channel %s[%s]", helper.GetIPAddressForRequest(r), len(comments), channel.Name, channel.ClaimID)
	}

	modifiedReactions, err := updateReactions(channel, args, commentIDs, comments)
	if err != nil {
		return errors.Err(err)
	}
	reply.Reactions = modifiedReactions
	return nil
}
func updateReactions(channel *model.Channel, args *commentapi.ReactArgs, commentIDs []interface{}, comments model.CommentSlice) (commentapi.Reactions, error) {
	var modifiedReactions = newReactions(strings.Split(args.CommentIDs, ","), &args.Type)
	err := db.WithTx(nil, func(tx boil.Transactor) error {
		if len(args.ClearTypes) > 0 {
			typeNames := util.StringSplitArg(args.ClearTypes, ",")
			reactionTypes, err := model.ReactionTypes(qm.WhereIn(model.ReactionTypeColumns.Name+" IN ?", typeNames...)).All(tx)
			if err != nil {
				return errors.Err(err)
			}
			if len(reactionTypes) > 0 {
				var typesToClear []interface{}
				for _, rt := range reactionTypes {
					typesToClear = append(typesToClear, rt.ID)
				}
				err = channel.Reactions(
					qm.Where(model.ReactionColumns.ChannelID+"=?", channel.ClaimID),
					qm.WhereIn(model.ReactionColumns.ReactionTypeID+" IN ?", typesToClear...),
					qm.WhereIn(model.ReactionColumns.CommentID+" IN ?", commentIDs...)).DeleteAll(tx)
				if err != nil {
					return errors.Err(err)
				}
			}
		}

		reactionType, err := model.ReactionTypes(model.ReactionTypeWhere.Name.EQ(args.Type)).One(tx)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
			reactionType = &model.ReactionType{Name: args.Type}
			err = reactionType.Insert(tx, boil.Infer())
		}
		if err != nil {
			return errors.Err(err)
		}
		if args.Remove {
			existingReactions, err := channel.Reactions(
				qm.WhereIn(model.ReactionColumns.CommentID+"=?", commentIDs...),
				qm.Where(model.ReactionColumns.ReactionTypeID+"=?", reactionType.ID),
				qm.Load("Comment")).All(tx)
			if err != nil {
				return errors.Err(err)
			}
			if len(existingReactions) == 0 {
				return api.StatusError{Err: errors.Err("there are no reactions for the claim(s) to remove"), Status: http.StatusBadRequest}
			}
			for _, r := range existingReactions {
				addTo(modifiedReactions[r.R.Comment.CommentID], args.Type)
			}
			err = existingReactions.DeleteAll(tx)
			return errors.Err(err)
		}
		for _, p := range comments {
			err = helper.AllowedToRespond(p.CommentID, channel.ClaimID)
			if err != nil {
				return err
			}
			newReaction := &model.Reaction{ChannelID: null.StringFrom(channel.ClaimID), CommentID: p.CommentID, ReactionTypeID: reactionType.ID, ClaimID: p.LbryClaimID, IsFlagged: len(comments) > 1}
			err := flags.CheckReaction(newReaction)
			if err != nil {
				return err
			}
			err = newReaction.Insert(tx, boil.Infer())
			if err != nil {
				if strings.Contains(err.Error(), "Duplicate entry") {
					return api.StatusError{Err: errors.Err("reaction already acknowledged!"), Status: http.StatusBadRequest}
				}
				return errors.Err(err)
			}
			addTo(modifiedReactions[p.CommentID], reactionType.Name)
		}
		return nil
	})
	if err != nil {
		return nil, errors.Err(err)
	}
	return modifiedReactions, nil
}
