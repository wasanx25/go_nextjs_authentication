import { PostsDialog } from '../uiParts/PostsDialog'

export const UserPosts = () => {
  async function handleRequest (text: string): Promise<void> {
    return await fetch('/api/post', {
      method: 'POST',
      body: JSON.stringify({text: text})
    }).then((response) => {
      if (response.ok) {
        return Promise.resolve()
      }

      return Promise.reject(`Failed request /api/post response status: ${response.status}`)
    }).catch((err) => {
      return Promise.reject(`Failed fetch function err: ${err}`)
    })
  }

  return (
    <>
      <PostsDialog handleSubmit={handleRequest} openModalButtonName={"投稿する"} headerTitle={"投稿する"}/>
    </>
  )
}
