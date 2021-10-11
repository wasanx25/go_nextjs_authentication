import React, { useEffect, useState } from 'react'
import { useToast, Button, Flex, ListItem, Text, UnorderedList, List, Box, Stack } from '@chakra-ui/react'

interface user {
  id: string
  username: string
}

export const FollowUserList = () => {
  const [users, setUsers] = useState([] as Array<user>)
  const toast = useToast()

  useEffect(() => {
    const f = async () => {
      const response = await fetch('/api/follow_users')
      const json = (await response.json())['users'] as Array<user>
      setUsers(json)
    }
    f()
  }, [])

  async function unfollowClickHandler (event: React.MouseEvent) {
    event.preventDefault()
    const id = event.currentTarget.getAttribute('data-user-id')
    await fetch(`/api/follow/${id}`, {method: 'DELETE'})
    .then((response) => {
      if (response.ok) {
        toast({
          title: 'unfollowed!',
          status: 'success',
          duration: 5000,
          isClosable: true
        })
      }
    })
  }

  return (
    <Flex>
      <List width='100%'>
        {users.map(s => {
          return (
            <ListItem p='0.5rem'>
              <Stack direction={['column', 'row']}>
                <Box>
                  <Text>{s.username}</Text>
                </Box>
                <Box>
                  <Button data-user-id={s.id} onClick={unfollowClickHandler}>unfollow</Button>
                </Box>
              </Stack>
            </ListItem>
          )
        })}
      </List>
    </Flex>
  )
}
