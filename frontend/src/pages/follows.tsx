import { AuthedTemplate } from '../components/templates/AuthedTemplate'
import { Box, Button, Flex, List, ListItem, Stack, Text, useToast } from '@chakra-ui/react'
import React, { useEffect, useState } from 'react'

interface user {
  id: string
  username: string
}

export default function Follows () {
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
    <AuthedTemplate>
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
    </AuthedTemplate>
  )
}
