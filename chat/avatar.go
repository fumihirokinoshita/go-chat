package main

import (
	"errors"
)

// ErrNoAvatarはAvatarインスタンスがアバターのURLを返すことができない場合に発生するerror
var ErrNoAvatarURL = errors.New("chat: アバターのURLを取得できません。")

// Avatarはユーザのプロフィール画像を表す型
type Avatar interface {
	// GetAvatarURLは指定されたクライアントのアバターのURLを返す
	// 問題がガッ生した場合はerrorを返す。特にURLを取得できなかった場合はErrNoAvatarURLを返す
	getAvatarURL(c *client) (string, error)
}
