import { AuthedTemplate } from '../components/templates/AuthedTemplate'
import { useEffect, useState } from 'react'
import { Box, Flex, List, ListItem, Text } from '@chakra-ui/react'

interface post {
  post_id: string
  user_id: string
  text: string
  posted_at: Date
}

export default function Index () {
  const [posts, setPosts] = useState([] as Array<post>)

  useEffect(() => {
    const f = async () => {
      const response = await fetch('/api/timeline')
      const json = (await response.json())['posts'] as Array<post>
      setPosts(json)
    }
    f()
  }, [])

  return (
    <AuthedTemplate>
      <Flex>
        <List width='100%' border='1px'>
          {posts.map(p => {
            return (
              <ListItem p='0.5rem'>
                <Box>
                  <Text fontSize='lg'>{p.text}</Text>
                  <Text as='time' fontSize='xs' color='gray.400'>{p.posted_at}</Text>
                </Box>
              </ListItem>
            )
          })}
        </List>
      </Flex>
    </AuthedTemplate>
  )
}
