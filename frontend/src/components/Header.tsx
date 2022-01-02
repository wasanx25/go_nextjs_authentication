import { Box, Button, Flex, Heading, Spacer, Stack, Text } from '@chakra-ui/react'
import Link from 'next/link'
import { DarkModeSwitch } from './uiParts/DarkModeSwitch'
import { useUser } from '@auth0/nextjs-auth0'
import { UserPosts } from './organisms/UserPosts'

export const Header = () => {
  const {user} = useUser()

  return (
    <Flex as="header" width='100%' p='1rem' borderBottom='1px' borderColor='gray.200'>
      <Box py='2'>
        <Heading size='md'>Application Title</Heading>
      </Box>
      <Spacer/>
      <Box>
        <Stack direction={['column', 'row']} spacing='0.5rem' alignItems='center'>
          <Text display='inline-block'>{user.name}</Text>

          <UserPosts/>

          <Button colorScheme='teal' variant='outline'>
            <Link href='/api/auth/logout'>
              <a>Logout</a>
            </Link>
          </Button>

          <DarkModeSwitch/>
        </Stack>
      </Box>
    </Flex>
  )
}
