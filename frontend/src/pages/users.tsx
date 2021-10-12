import { UserList } from '../components/UserList'
import { AuthedTemplate } from '../components/AuthedTemplate'

export default function Users () {
  return (
    <AuthedTemplate>
      <UserList/>
    </AuthedTemplate>
  )
}
