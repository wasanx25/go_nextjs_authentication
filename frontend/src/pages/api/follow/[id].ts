import { getSession } from '@auth0/nextjs-auth0'
import { NextApiRequest, NextApiResponse } from 'next'

export default async function no_follow_users (req: NextApiRequest, res: NextApiResponse) {
  const session = getSession(req, res)
  const token = session.idToken

  const {
    query: { id },
    method,
  } = req

  switch (method) {
    case 'POST':
    case 'DELETE':
      const response = await fetch(`http://localhost:1323/follow/${id}`, {
        method: method,
        headers: {
          'Authorization': `Bearer ${token}`
        }
      })

      res.status(response.status).end()
      break
    default:
      res.setHeader('Allow', ['POST', 'DELETE'])
      res.status(405).end(`Method ${method} Not Allowed`)
  }
}
