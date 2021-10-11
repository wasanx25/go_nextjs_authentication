import { Box, Button, Flex, Heading, Spacer, Stack, Text } from '@chakra-ui/react'
import { DarkModeSwitch } from './DarkModeSwitch'
import { PostsDialog } from './PostsDialog'

interface HeaderProps {
  username: string
}

export const Header = (props: HeaderProps) => {
  return (
    <Flex width='100%' p='1rem' borderBottom='1px' borderColor="gray.200">
      <Box py='2'>
        <Heading size='md'>Application Title</Heading>
      </Box>
      <Spacer/>
      <Box>
        <Stack direction={['column', 'row']} spacing='0.5rem' alignItems='center'>
          <Text display='inline-block'>{props.username}</Text>

          <PostsDialog/>

          <Button colorScheme='teal' variant='outline'>
            <a href='/api/auth/logout'>Logout</a>
          </Button>

          <DarkModeSwitch/>
        </Stack>
      </Box>
    </Flex>
  )
}
