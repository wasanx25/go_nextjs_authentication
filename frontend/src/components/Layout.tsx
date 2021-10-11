import { Box, Button, Flex, Heading, Spacer, Stack, Text } from '@chakra-ui/react'
import { DarkModeSwitch } from './DarkModeSwitch'
import { PostsDialog } from './PostsDialog'
import { ReactNode } from 'react'
import { Container } from './Container'
import { Header } from './Header'
import { Sidebar } from './Sidebar'
import { Main } from './Main'
import { Timeline } from './Timeline'
import { Footer } from './Footer'
import { useUser } from '@auth0/nextjs-auth0'
import { UserList } from './UserList'

interface LayoutProps {
  children: ReactNode
}

export const Layout = (props: LayoutProps) => {
  const {user, error, isLoading} = useUser()

  if (isLoading) return <div>Loading...</div>
  if (error) return <div>{error.message}</div>

  if (user) {
    return (
      <Container>
        <Header username={user.name}/>
        <Flex width='100%'>
          <Box width='20%'>
            <Sidebar/>
          </Box>
          <Box width='60%'>
            <Main>
              {props.children}
            </Main>
          </Box>
        </Flex>

        <Footer>
          <Text>Footer</Text>
        </Footer>
      </Container>
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
