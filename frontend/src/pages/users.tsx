import { UserList } from '../components/UserList'
import { AuthedTemplate } from '../components/templates/AuthedTemplate'

export default function Users () {
  return (
    <AuthedTemplate>
      <UserList/>
    </AuthedTemplate>
  )
}
