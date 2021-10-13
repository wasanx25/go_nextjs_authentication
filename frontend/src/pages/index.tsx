import { Timeline } from '../components/Timeline'
import { AuthedTemplate } from '../components/templates/AuthedTemplate'

export default function Index () {
  return (
    <AuthedTemplate>
      <Timeline/>
    </AuthedTemplate>
  )
}
