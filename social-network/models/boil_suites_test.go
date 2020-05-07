// Code generated by SQLBoiler 4.1.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Comments", testComments)
	t.Run("CommentLikes", testCommentLikes)
	t.Run("Friendships", testFriendships)
	t.Run("Images", testImages)
	t.Run("ImageVariations", testImageVariations)
	t.Run("Likes", testLikes)
	t.Run("Posts", testPosts)
	t.Run("Users", testUsers)
}

func TestDelete(t *testing.T) {
	t.Run("Comments", testCommentsDelete)
	t.Run("CommentLikes", testCommentLikesDelete)
	t.Run("Friendships", testFriendshipsDelete)
	t.Run("Images", testImagesDelete)
	t.Run("ImageVariations", testImageVariationsDelete)
	t.Run("Likes", testLikesDelete)
	t.Run("Posts", testPostsDelete)
	t.Run("Users", testUsersDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Comments", testCommentsQueryDeleteAll)
	t.Run("CommentLikes", testCommentLikesQueryDeleteAll)
	t.Run("Friendships", testFriendshipsQueryDeleteAll)
	t.Run("Images", testImagesQueryDeleteAll)
	t.Run("ImageVariations", testImageVariationsQueryDeleteAll)
	t.Run("Likes", testLikesQueryDeleteAll)
	t.Run("Posts", testPostsQueryDeleteAll)
	t.Run("Users", testUsersQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Comments", testCommentsSliceDeleteAll)
	t.Run("CommentLikes", testCommentLikesSliceDeleteAll)
	t.Run("Friendships", testFriendshipsSliceDeleteAll)
	t.Run("Images", testImagesSliceDeleteAll)
	t.Run("ImageVariations", testImageVariationsSliceDeleteAll)
	t.Run("Likes", testLikesSliceDeleteAll)
	t.Run("Posts", testPostsSliceDeleteAll)
	t.Run("Users", testUsersSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Comments", testCommentsExists)
	t.Run("CommentLikes", testCommentLikesExists)
	t.Run("Friendships", testFriendshipsExists)
	t.Run("Images", testImagesExists)
	t.Run("ImageVariations", testImageVariationsExists)
	t.Run("Likes", testLikesExists)
	t.Run("Posts", testPostsExists)
	t.Run("Users", testUsersExists)
}

func TestFind(t *testing.T) {
	t.Run("Comments", testCommentsFind)
	t.Run("CommentLikes", testCommentLikesFind)
	t.Run("Friendships", testFriendshipsFind)
	t.Run("Images", testImagesFind)
	t.Run("ImageVariations", testImageVariationsFind)
	t.Run("Likes", testLikesFind)
	t.Run("Posts", testPostsFind)
	t.Run("Users", testUsersFind)
}

func TestBind(t *testing.T) {
	t.Run("Comments", testCommentsBind)
	t.Run("CommentLikes", testCommentLikesBind)
	t.Run("Friendships", testFriendshipsBind)
	t.Run("Images", testImagesBind)
	t.Run("ImageVariations", testImageVariationsBind)
	t.Run("Likes", testLikesBind)
	t.Run("Posts", testPostsBind)
	t.Run("Users", testUsersBind)
}

func TestOne(t *testing.T) {
	t.Run("Comments", testCommentsOne)
	t.Run("CommentLikes", testCommentLikesOne)
	t.Run("Friendships", testFriendshipsOne)
	t.Run("Images", testImagesOne)
	t.Run("ImageVariations", testImageVariationsOne)
	t.Run("Likes", testLikesOne)
	t.Run("Posts", testPostsOne)
	t.Run("Users", testUsersOne)
}

func TestAll(t *testing.T) {
	t.Run("Comments", testCommentsAll)
	t.Run("CommentLikes", testCommentLikesAll)
	t.Run("Friendships", testFriendshipsAll)
	t.Run("Images", testImagesAll)
	t.Run("ImageVariations", testImageVariationsAll)
	t.Run("Likes", testLikesAll)
	t.Run("Posts", testPostsAll)
	t.Run("Users", testUsersAll)
}

func TestCount(t *testing.T) {
	t.Run("Comments", testCommentsCount)
	t.Run("CommentLikes", testCommentLikesCount)
	t.Run("Friendships", testFriendshipsCount)
	t.Run("Images", testImagesCount)
	t.Run("ImageVariations", testImageVariationsCount)
	t.Run("Likes", testLikesCount)
	t.Run("Posts", testPostsCount)
	t.Run("Users", testUsersCount)
}

func TestHooks(t *testing.T) {
	t.Run("Comments", testCommentsHooks)
	t.Run("CommentLikes", testCommentLikesHooks)
	t.Run("Friendships", testFriendshipsHooks)
	t.Run("Images", testImagesHooks)
	t.Run("ImageVariations", testImageVariationsHooks)
	t.Run("Likes", testLikesHooks)
	t.Run("Posts", testPostsHooks)
	t.Run("Users", testUsersHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Comments", testCommentsInsert)
	t.Run("Comments", testCommentsInsertWhitelist)
	t.Run("CommentLikes", testCommentLikesInsert)
	t.Run("CommentLikes", testCommentLikesInsertWhitelist)
	t.Run("Friendships", testFriendshipsInsert)
	t.Run("Friendships", testFriendshipsInsertWhitelist)
	t.Run("Images", testImagesInsert)
	t.Run("Images", testImagesInsertWhitelist)
	t.Run("ImageVariations", testImageVariationsInsert)
	t.Run("ImageVariations", testImageVariationsInsertWhitelist)
	t.Run("Likes", testLikesInsert)
	t.Run("Likes", testLikesInsertWhitelist)
	t.Run("Posts", testPostsInsert)
	t.Run("Posts", testPostsInsertWhitelist)
	t.Run("Users", testUsersInsert)
	t.Run("Users", testUsersInsertWhitelist)
}

// TestToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestToOne(t *testing.T) {
	t.Run("CommentToPostUsingPost", testCommentToOnePostUsingPost)
	t.Run("CommentToUserUsingUser", testCommentToOneUserUsingUser)
	t.Run("CommentLikeToCommentUsingComment", testCommentLikeToOneCommentUsingComment)
	t.Run("CommentLikeToUserUsingUser", testCommentLikeToOneUserUsingUser)
	t.Run("ImageToPostUsingPost", testImageToOnePostUsingPost)
	t.Run("ImageVariationToImageUsingImage", testImageVariationToOneImageUsingImage)
	t.Run("LikeToPostUsingPost", testLikeToOnePostUsingPost)
	t.Run("LikeToUserUsingUser", testLikeToOneUserUsingUser)
	t.Run("PostToUserUsingUser", testPostToOneUserUsingUser)
}

// TestOneToOne tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOne(t *testing.T) {}

// TestToMany tests cannot be run in parallel
// or deadlocks can occur.
func TestToMany(t *testing.T) {
	t.Run("CommentToCommentLikes", testCommentToManyCommentLikes)
	t.Run("FriendshipToUsers", testFriendshipToManyUsers)
	t.Run("ImageToImageVariations", testImageToManyImageVariations)
	t.Run("PostToComments", testPostToManyComments)
	t.Run("PostToImages", testPostToManyImages)
	t.Run("PostToLikes", testPostToManyLikes)
	t.Run("UserToComments", testUserToManyComments)
	t.Run("UserToCommentLikes", testUserToManyCommentLikes)
	t.Run("UserToLikes", testUserToManyLikes)
	t.Run("UserToPosts", testUserToManyPosts)
	t.Run("UserToFriendships", testUserToManyFriendships)
}

// TestToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneSet(t *testing.T) {
	t.Run("CommentToPostUsingComments", testCommentToOneSetOpPostUsingPost)
	t.Run("CommentToUserUsingComments", testCommentToOneSetOpUserUsingUser)
	t.Run("CommentLikeToCommentUsingCommentLikes", testCommentLikeToOneSetOpCommentUsingComment)
	t.Run("CommentLikeToUserUsingCommentLikes", testCommentLikeToOneSetOpUserUsingUser)
	t.Run("ImageToPostUsingImages", testImageToOneSetOpPostUsingPost)
	t.Run("ImageVariationToImageUsingImageVariations", testImageVariationToOneSetOpImageUsingImage)
	t.Run("LikeToPostUsingLikes", testLikeToOneSetOpPostUsingPost)
	t.Run("LikeToUserUsingLikes", testLikeToOneSetOpUserUsingUser)
	t.Run("PostToUserUsingPosts", testPostToOneSetOpUserUsingUser)
}

// TestToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToOneRemove(t *testing.T) {}

// TestOneToOneSet tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneSet(t *testing.T) {}

// TestOneToOneRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestOneToOneRemove(t *testing.T) {}

// TestToManyAdd tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyAdd(t *testing.T) {
	t.Run("CommentToCommentLikes", testCommentToManyAddOpCommentLikes)
	t.Run("FriendshipToUsers", testFriendshipToManyAddOpUsers)
	t.Run("ImageToImageVariations", testImageToManyAddOpImageVariations)
	t.Run("PostToComments", testPostToManyAddOpComments)
	t.Run("PostToImages", testPostToManyAddOpImages)
	t.Run("PostToLikes", testPostToManyAddOpLikes)
	t.Run("UserToComments", testUserToManyAddOpComments)
	t.Run("UserToCommentLikes", testUserToManyAddOpCommentLikes)
	t.Run("UserToLikes", testUserToManyAddOpLikes)
	t.Run("UserToPosts", testUserToManyAddOpPosts)
	t.Run("UserToFriendships", testUserToManyAddOpFriendships)
}

// TestToManySet tests cannot be run in parallel
// or deadlocks can occur.
func TestToManySet(t *testing.T) {
	t.Run("FriendshipToUsers", testFriendshipToManySetOpUsers)
	t.Run("UserToFriendships", testUserToManySetOpFriendships)
}

// TestToManyRemove tests cannot be run in parallel
// or deadlocks can occur.
func TestToManyRemove(t *testing.T) {
	t.Run("FriendshipToUsers", testFriendshipToManyRemoveOpUsers)
	t.Run("UserToFriendships", testUserToManyRemoveOpFriendships)
}

func TestReload(t *testing.T) {
	t.Run("Comments", testCommentsReload)
	t.Run("CommentLikes", testCommentLikesReload)
	t.Run("Friendships", testFriendshipsReload)
	t.Run("Images", testImagesReload)
	t.Run("ImageVariations", testImageVariationsReload)
	t.Run("Likes", testLikesReload)
	t.Run("Posts", testPostsReload)
	t.Run("Users", testUsersReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Comments", testCommentsReloadAll)
	t.Run("CommentLikes", testCommentLikesReloadAll)
	t.Run("Friendships", testFriendshipsReloadAll)
	t.Run("Images", testImagesReloadAll)
	t.Run("ImageVariations", testImageVariationsReloadAll)
	t.Run("Likes", testLikesReloadAll)
	t.Run("Posts", testPostsReloadAll)
	t.Run("Users", testUsersReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Comments", testCommentsSelect)
	t.Run("CommentLikes", testCommentLikesSelect)
	t.Run("Friendships", testFriendshipsSelect)
	t.Run("Images", testImagesSelect)
	t.Run("ImageVariations", testImageVariationsSelect)
	t.Run("Likes", testLikesSelect)
	t.Run("Posts", testPostsSelect)
	t.Run("Users", testUsersSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Comments", testCommentsUpdate)
	t.Run("CommentLikes", testCommentLikesUpdate)
	t.Run("Friendships", testFriendshipsUpdate)
	t.Run("Images", testImagesUpdate)
	t.Run("ImageVariations", testImageVariationsUpdate)
	t.Run("Likes", testLikesUpdate)
	t.Run("Posts", testPostsUpdate)
	t.Run("Users", testUsersUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Comments", testCommentsSliceUpdateAll)
	t.Run("CommentLikes", testCommentLikesSliceUpdateAll)
	t.Run("Friendships", testFriendshipsSliceUpdateAll)
	t.Run("Images", testImagesSliceUpdateAll)
	t.Run("ImageVariations", testImageVariationsSliceUpdateAll)
	t.Run("Likes", testLikesSliceUpdateAll)
	t.Run("Posts", testPostsSliceUpdateAll)
	t.Run("Users", testUsersSliceUpdateAll)
}
