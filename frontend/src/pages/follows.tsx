import { FollowUserList } from '../components/FollowUserList'
import { AuthedTemplate } from '../components/templates/AuthedTemplate'

export default function Follows () {
  return (
    <AuthedTemplate>
      <FollowUserList/>
    </AuthedTemplate>
  )
}
