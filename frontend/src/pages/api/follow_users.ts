import { getSession } from '@auth0/nextjs-auth0'
import { NextApiRequest, NextApiResponse } from 'next'

export default async function no_follow_users (req: NextApiRequest, res: NextApiResponse) {
  const session = getSession(req, res)
  const token = session.idToken

  const response = await fetch('http://localhost:1323/follow_users', {
    headers: {
      'Authorization': `Bearer ${token}`
    }
  })
  const json = await response.json()

  res.send(JSON.stringify(json, null, 2))
}
