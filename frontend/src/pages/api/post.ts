import { getSession } from '@auth0/nextjs-auth0'
import { NextApiRequest, NextApiResponse } from 'next'

export default async function post (req: NextApiRequest, res: NextApiResponse) {
  const session = getSession(req, res)
  const token = session.idToken

  await fetch('http://localhost:1323/post', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`
    },
    body: req.body
  }).then((response) => {
    res.status(response.status)
      .send(response.body)
  }).catch((reason) => {
    console.error(reason)
  })
}
