import { Button } from '@chakra-ui/react'
import React, { ReactNode } from 'react'
import { Container } from './Container'
import { useUser } from '@auth0/nextjs-auth0'
import { Layout } from './Layout'

export const AuthedTemplate: React.FC<{children: ReactNode}> = ({children}) => {
  const {user, error, isLoading} = useUser()

  if (isLoading) return <div>Loading...</div>
  if (error) return <div>{error.message}</div>

  if (user) {
    return (
      <Layout>
        {children}
      </Layout>
    )
  }

  return (
    <Container height='100vh' pt='4rem'>
      <Button colorScheme="teal" variant="solid" size='lg'>
        <a href='/api/auth/login'>Login</a>
      </Button>
    </Container>
  )
}
