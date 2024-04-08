import { defineStore } from 'pinia';

export default defineStore("feed", {
  state: () => ({
    posts: [],
    userPosts: [],
    groupPost: {}
  }),
  actions: {
    addPost(post) {
      if (post.group_id!="00000000-0000-0000-0000-000000000000"){
        console.log(post, "dslmdsldmslmdlsmldmsldmslmdlmsl");
        this.groupPost[post.group_id].unshift(post)
        return 
      }
      this.posts.unshift(post);
      this.userPosts.unshift(post)
    },
    async getUserFeed() {
      await fetch("/api/getFeed")
        .then(async (response) => {
          const data = await response.json()
          this.posts = data.body
          for (let i = 0; i < this.posts.length; i++) {
            if (this.posts[i].userOwnerNickname === useAuthUser().value.nickname) {
              this.userPosts.push(this.posts[i]);
            }
          }
        })
        .catch((error) => console.error(error))
    },
    addComment(comment) {
      if (comment.postPrivacy== "group"){
        let groupPostS = this.groupPost[comment.group_id]
        for (let i = 0; i < groupPostS.length; i++) {
          if (groupPostS[i].id === comment.post_id) {
            groupPostS[i].comments.push(comment);
            return
          }
        }
      }

      for (let i = 0; i < this.posts.length; i++) {
        if (this.posts[i].id === comment.post_id) {

          this.posts[i].comments.push(comment);
          return
        }

      }
    },
    getUserPosts(nickname) {

      let userPosts = []
      for (let post of this.posts) {

        if (post.userOwnerNickname === nickname) {
          userPosts.push(post)
        }
      }
      return userPosts
    },
    flushAllPosts() {
      this.userPosts = []
      this.posts = []
    },
    async getGroupPosts(groupId) {
      console.log({ groupId });
      let response = await fetch(`/api/post/groups`, {
        headers: { groupId }
      })
        .then(async (response) => {
          const data = await response.json()
          return data
        })
        .catch((error) => console.error(error))
        this.groupPost[groupId] = response.body.data
    },

  }
});