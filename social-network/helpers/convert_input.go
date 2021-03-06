// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package helpers

import (
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/social-network/graphql_models"
	"github.com/web-ridge/gqlgen-sqlboiler-examples/social-network/models"
	"github.com/web-ridge/utils-go/boilergql"
)

func CommentCreateInputsToBoiler(am []*graphql_models.CommentCreateInput) []*models.Comment {
	ar := make([]*models.Comment, len(am))
	for i, m := range am {
		ar[i] = CommentCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func CommentCreateInputToBoiler(
	m *graphql_models.CommentCreateInput,
) *models.Comment {
	if m == nil {
		return nil
	}
	r := &models.Comment{
		Content: m.Content,
		PostID:  boilergql.IDToBoiler(m.PostID),
	}
	return r
}

func CommentCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.CommentCreateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "content":
			modelM[models.CommentColumns.Content] = m.Content
		case "postId":
			modelM[models.CommentColumns.PostID] = boilergql.IDToBoiler(m.PostID)
		}
	}
	return modelM
}

func CommentCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "content":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentColumns.Content)
		case "postId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentColumns.PostID)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func CommentLikeCreateInputsToBoiler(am []*graphql_models.CommentLikeCreateInput) []*models.CommentLike {
	ar := make([]*models.CommentLike, len(am))
	for i, m := range am {
		ar[i] = CommentLikeCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func CommentLikeCreateInputToBoiler(
	m *graphql_models.CommentLikeCreateInput,
) *models.CommentLike {
	if m == nil {
		return nil
	}
	r := &models.CommentLike{
		CommentID: boilergql.IDToBoiler(m.CommentID),
		LikeType:  m.LikeType,
		CreatedAt: boilergql.PointerIntToNullDotTime(m.CreatedAt),
	}
	return r
}

func CommentLikeCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.CommentLikeCreateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "commentId":
			modelM[models.CommentLikeColumns.CommentID] = boilergql.IDToBoiler(m.CommentID)
		case "likeType":
			modelM[models.CommentLikeColumns.LikeType] = m.LikeType
		case "createdAt":
			modelM[models.CommentLikeColumns.CreatedAt] = boilergql.PointerIntToNullDotTime(m.CreatedAt)
		}
	}
	return modelM
}

func CommentLikeCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "commentId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentLikeColumns.CommentID)
		case "likeType":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentLikeColumns.LikeType)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentLikeColumns.CreatedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func CommentLikeUpdateInputsToBoiler(am []*graphql_models.CommentLikeUpdateInput) []*models.CommentLike {
	ar := make([]*models.CommentLike, len(am))
	for i, m := range am {
		ar[i] = CommentLikeUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func CommentLikeUpdateInputToBoiler(
	m *graphql_models.CommentLikeUpdateInput,
) *models.CommentLike {
	if m == nil {
		return nil
	}
	r := &models.CommentLike{
		CommentID: boilergql.IDToBoiler(boilergql.PointerStringToString(m.CommentID)),
		LikeType:  boilergql.PointerStringToString(m.LikeType),
		CreatedAt: boilergql.PointerIntToNullDotTime(m.CreatedAt),
	}
	return r
}

func CommentLikeUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.CommentLikeUpdateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "commentId":
			modelM[models.CommentLikeColumns.CommentID] = boilergql.IDToBoiler(boilergql.PointerStringToString(m.CommentID))
		case "likeType":
			modelM[models.CommentLikeColumns.LikeType] = boilergql.PointerStringToString(m.LikeType)
		case "createdAt":
			modelM[models.CommentLikeColumns.CreatedAt] = boilergql.PointerIntToNullDotTime(m.CreatedAt)
		}
	}
	return modelM
}

func CommentLikeUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "commentId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentLikeColumns.CommentID)
		case "likeType":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentLikeColumns.LikeType)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentLikeColumns.CreatedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func CommentUpdateInputsToBoiler(am []*graphql_models.CommentUpdateInput) []*models.Comment {
	ar := make([]*models.Comment, len(am))
	for i, m := range am {
		ar[i] = CommentUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func CommentUpdateInputToBoiler(
	m *graphql_models.CommentUpdateInput,
) *models.Comment {
	if m == nil {
		return nil
	}
	r := &models.Comment{
		Content: boilergql.PointerStringToString(m.Content),
		PostID:  boilergql.IDToBoiler(boilergql.PointerStringToString(m.PostID)),
	}
	return r
}

func CommentUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.CommentUpdateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "content":
			modelM[models.CommentColumns.Content] = boilergql.PointerStringToString(m.Content)
		case "postId":
			modelM[models.CommentColumns.PostID] = boilergql.IDToBoiler(boilergql.PointerStringToString(m.PostID))
		}
	}
	return modelM
}

func CommentUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "content":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentColumns.Content)
		case "postId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.CommentColumns.PostID)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func FriendshipCreateInputsToBoiler(am []*graphql_models.FriendshipCreateInput) []*models.Friendship {
	ar := make([]*models.Friendship, len(am))
	for i, m := range am {
		ar[i] = FriendshipCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func FriendshipCreateInputToBoiler(
	m *graphql_models.FriendshipCreateInput,
) *models.Friendship {
	if m == nil {
		return nil
	}
	r := &models.Friendship{
		CreatedAt: boilergql.PointerIntToNullDotTime(m.CreatedAt),
	}
	return r
}

func FriendshipCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.FriendshipCreateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "createdAt":
			modelM[models.FriendshipColumns.CreatedAt] = boilergql.PointerIntToNullDotTime(m.CreatedAt)
		}
	}
	return modelM
}

func FriendshipCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FriendshipColumns.CreatedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func FriendshipUpdateInputsToBoiler(am []*graphql_models.FriendshipUpdateInput) []*models.Friendship {
	ar := make([]*models.Friendship, len(am))
	for i, m := range am {
		ar[i] = FriendshipUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func FriendshipUpdateInputToBoiler(
	m *graphql_models.FriendshipUpdateInput,
) *models.Friendship {
	if m == nil {
		return nil
	}
	r := &models.Friendship{
		CreatedAt: boilergql.PointerIntToNullDotTime(m.CreatedAt),
	}
	return r
}

func FriendshipUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.FriendshipUpdateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "createdAt":
			modelM[models.FriendshipColumns.CreatedAt] = boilergql.PointerIntToNullDotTime(m.CreatedAt)
		}
	}
	return modelM
}

func FriendshipUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.FriendshipColumns.CreatedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func ImageCreateInputsToBoiler(am []*graphql_models.ImageCreateInput) []*models.Image {
	ar := make([]*models.Image, len(am))
	for i, m := range am {
		ar[i] = ImageCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func ImageCreateInputToBoiler(
	m *graphql_models.ImageCreateInput,
) *models.Image {
	if m == nil {
		return nil
	}
	r := &models.Image{
		PostID:      boilergql.IDToBoiler(m.PostID),
		Views:       boilergql.PointerIntToNullDotInt(m.Views),
		OriginalURL: boilergql.PointerStringToNullDotString(m.OriginalURL),
	}
	return r
}

func ImageCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.ImageCreateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "postId":
			modelM[models.ImageColumns.PostID] = boilergql.IDToBoiler(m.PostID)
		case "views":
			modelM[models.ImageColumns.Views] = boilergql.PointerIntToNullDotInt(m.Views)
		case "originalUrl":
			modelM[models.ImageColumns.OriginalURL] = boilergql.PointerStringToNullDotString(m.OriginalURL)
		}
	}
	return modelM
}

func ImageCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "postId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageColumns.PostID)
		case "views":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageColumns.Views)
		case "originalUrl":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageColumns.OriginalURL)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func ImageUpdateInputsToBoiler(am []*graphql_models.ImageUpdateInput) []*models.Image {
	ar := make([]*models.Image, len(am))
	for i, m := range am {
		ar[i] = ImageUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func ImageUpdateInputToBoiler(
	m *graphql_models.ImageUpdateInput,
) *models.Image {
	if m == nil {
		return nil
	}
	r := &models.Image{
		PostID:      boilergql.IDToBoiler(boilergql.PointerStringToString(m.PostID)),
		Views:       boilergql.PointerIntToNullDotInt(m.Views),
		OriginalURL: boilergql.PointerStringToNullDotString(m.OriginalURL),
	}
	return r
}

func ImageUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.ImageUpdateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "postId":
			modelM[models.ImageColumns.PostID] = boilergql.IDToBoiler(boilergql.PointerStringToString(m.PostID))
		case "views":
			modelM[models.ImageColumns.Views] = boilergql.PointerIntToNullDotInt(m.Views)
		case "originalUrl":
			modelM[models.ImageColumns.OriginalURL] = boilergql.PointerStringToNullDotString(m.OriginalURL)
		}
	}
	return modelM
}

func ImageUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "postId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageColumns.PostID)
		case "views":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageColumns.Views)
		case "originalUrl":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageColumns.OriginalURL)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func ImageVariationCreateInputsToBoiler(am []*graphql_models.ImageVariationCreateInput) []*models.ImageVariation {
	ar := make([]*models.ImageVariation, len(am))
	for i, m := range am {
		ar[i] = ImageVariationCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func ImageVariationCreateInputToBoiler(
	m *graphql_models.ImageVariationCreateInput,
) *models.ImageVariation {
	if m == nil {
		return nil
	}
	r := &models.ImageVariation{
		ImageID: boilergql.IDToBoiler(m.ImageID),
	}
	return r
}

func ImageVariationCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.ImageVariationCreateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "imageId":
			modelM[models.ImageVariationColumns.ImageID] = boilergql.IDToBoiler(m.ImageID)
		}
	}
	return modelM
}

func ImageVariationCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "imageId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageVariationColumns.ImageID)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func ImageVariationUpdateInputsToBoiler(am []*graphql_models.ImageVariationUpdateInput) []*models.ImageVariation {
	ar := make([]*models.ImageVariation, len(am))
	for i, m := range am {
		ar[i] = ImageVariationUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func ImageVariationUpdateInputToBoiler(
	m *graphql_models.ImageVariationUpdateInput,
) *models.ImageVariation {
	if m == nil {
		return nil
	}
	r := &models.ImageVariation{
		ImageID: boilergql.IDToBoiler(boilergql.PointerStringToString(m.ImageID)),
	}
	return r
}

func ImageVariationUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.ImageVariationUpdateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "imageId":
			modelM[models.ImageVariationColumns.ImageID] = boilergql.IDToBoiler(boilergql.PointerStringToString(m.ImageID))
		}
	}
	return modelM
}

func ImageVariationUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "imageId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.ImageVariationColumns.ImageID)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func LikeCreateInputsToBoiler(am []*graphql_models.LikeCreateInput) []*models.Like {
	ar := make([]*models.Like, len(am))
	for i, m := range am {
		ar[i] = LikeCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func LikeCreateInputToBoiler(
	m *graphql_models.LikeCreateInput,
) *models.Like {
	if m == nil {
		return nil
	}
	r := &models.Like{
		PostID:    boilergql.IDToBoiler(m.PostID),
		LikeType:  m.LikeType,
		CreatedAt: boilergql.PointerIntToNullDotTime(m.CreatedAt),
	}
	return r
}

func LikeCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.LikeCreateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "postId":
			modelM[models.LikeColumns.PostID] = boilergql.IDToBoiler(m.PostID)
		case "likeType":
			modelM[models.LikeColumns.LikeType] = m.LikeType
		case "createdAt":
			modelM[models.LikeColumns.CreatedAt] = boilergql.PointerIntToNullDotTime(m.CreatedAt)
		}
	}
	return modelM
}

func LikeCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "postId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LikeColumns.PostID)
		case "likeType":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LikeColumns.LikeType)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LikeColumns.CreatedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func LikeUpdateInputsToBoiler(am []*graphql_models.LikeUpdateInput) []*models.Like {
	ar := make([]*models.Like, len(am))
	for i, m := range am {
		ar[i] = LikeUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func LikeUpdateInputToBoiler(
	m *graphql_models.LikeUpdateInput,
) *models.Like {
	if m == nil {
		return nil
	}
	r := &models.Like{
		PostID:    boilergql.IDToBoiler(boilergql.PointerStringToString(m.PostID)),
		LikeType:  boilergql.PointerStringToString(m.LikeType),
		CreatedAt: boilergql.PointerIntToNullDotTime(m.CreatedAt),
	}
	return r
}

func LikeUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.LikeUpdateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "postId":
			modelM[models.LikeColumns.PostID] = boilergql.IDToBoiler(boilergql.PointerStringToString(m.PostID))
		case "likeType":
			modelM[models.LikeColumns.LikeType] = boilergql.PointerStringToString(m.LikeType)
		case "createdAt":
			modelM[models.LikeColumns.CreatedAt] = boilergql.PointerIntToNullDotTime(m.CreatedAt)
		}
	}
	return modelM
}

func LikeUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "postId":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LikeColumns.PostID)
		case "likeType":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LikeColumns.LikeType)
		case "createdAt":
			columnsWhichAreSet = append(columnsWhichAreSet, models.LikeColumns.CreatedAt)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func PostCreateInputsToBoiler(am []*graphql_models.PostCreateInput) []*models.Post {
	ar := make([]*models.Post, len(am))
	for i, m := range am {
		ar[i] = PostCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func PostCreateInputToBoiler(
	m *graphql_models.PostCreateInput,
) *models.Post {
	if m == nil {
		return nil
	}
	r := &models.Post{
		Content: m.Content,
	}
	return r
}

func PostCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.PostCreateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "content":
			modelM[models.PostColumns.Content] = m.Content
		}
	}
	return modelM
}

func PostCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "content":
			columnsWhichAreSet = append(columnsWhichAreSet, models.PostColumns.Content)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func PostUpdateInputsToBoiler(am []*graphql_models.PostUpdateInput) []*models.Post {
	ar := make([]*models.Post, len(am))
	for i, m := range am {
		ar[i] = PostUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func PostUpdateInputToBoiler(
	m *graphql_models.PostUpdateInput,
) *models.Post {
	if m == nil {
		return nil
	}
	r := &models.Post{
		Content: boilergql.PointerStringToString(m.Content),
	}
	return r
}

func PostUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.PostUpdateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "content":
			modelM[models.PostColumns.Content] = boilergql.PointerStringToString(m.Content)
		}
	}
	return modelM
}

func PostUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "content":
			columnsWhichAreSet = append(columnsWhichAreSet, models.PostColumns.Content)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func UserCreateInputsToBoiler(am []*graphql_models.UserCreateInput) []*models.User {
	ar := make([]*models.User, len(am))
	for i, m := range am {
		ar[i] = UserCreateInputToBoiler(
			m,
		)
	}
	return ar
}

func UserCreateInputToBoiler(
	m *graphql_models.UserCreateInput,
) *models.User {
	if m == nil {
		return nil
	}
	r := &models.User{
		FirstName: m.FirstName,
		LastName:  m.LastName,
		Email:     m.Email,
	}
	return r
}

func UserCreateInputToModelM(
	input map[string]interface{},
	m graphql_models.UserCreateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "firstName":
			modelM[models.UserColumns.FirstName] = m.FirstName
		case "lastName":
			modelM[models.UserColumns.LastName] = m.LastName
		case "email":
			modelM[models.UserColumns.Email] = m.Email
		}
	}
	return modelM
}

func UserCreateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "firstName":
			columnsWhichAreSet = append(columnsWhichAreSet, models.UserColumns.FirstName)
		case "lastName":
			columnsWhichAreSet = append(columnsWhichAreSet, models.UserColumns.LastName)
		case "email":
			columnsWhichAreSet = append(columnsWhichAreSet, models.UserColumns.Email)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}

func UserUpdateInputsToBoiler(am []*graphql_models.UserUpdateInput) []*models.User {
	ar := make([]*models.User, len(am))
	for i, m := range am {
		ar[i] = UserUpdateInputToBoiler(
			m,
		)
	}
	return ar
}

func UserUpdateInputToBoiler(
	m *graphql_models.UserUpdateInput,
) *models.User {
	if m == nil {
		return nil
	}
	r := &models.User{
		FirstName: boilergql.PointerStringToString(m.FirstName),
		LastName:  boilergql.PointerStringToString(m.LastName),
		Email:     boilergql.PointerStringToString(m.Email),
	}
	return r
}

func UserUpdateInputToModelM(
	input map[string]interface{},
	m graphql_models.UserUpdateInput,
) models.M {
	modelM := models.M{}
	for key, _ := range input {
		switch key {
		case "firstName":
			modelM[models.UserColumns.FirstName] = boilergql.PointerStringToString(m.FirstName)
		case "lastName":
			modelM[models.UserColumns.LastName] = boilergql.PointerStringToString(m.LastName)
		case "email":
			modelM[models.UserColumns.Email] = boilergql.PointerStringToString(m.Email)
		}
	}
	return modelM
}

func UserUpdateInputToBoilerWhitelist(input map[string]interface{}, extraColumns ...string) boil.Columns {
	columnsWhichAreSet := []string{}
	for key, _ := range input {
		switch key {
		case "firstName":
			columnsWhichAreSet = append(columnsWhichAreSet, models.UserColumns.FirstName)
		case "lastName":
			columnsWhichAreSet = append(columnsWhichAreSet, models.UserColumns.LastName)
		case "email":
			columnsWhichAreSet = append(columnsWhichAreSet, models.UserColumns.Email)
		}
	}
	columnsWhichAreSet = append(columnsWhichAreSet, extraColumns...)
	return boil.Whitelist(columnsWhichAreSet...)
}
