import { AuthedTemplate } from '../components/templates/AuthedTemplate'
import { Box, Button, Flex, List, ListItem, Stack, Text, useToast } from '@chakra-ui/react'
import React, { useEffect, useState } from 'react'

interface user {
  id: string
  username: string
}

export default function Users () {
  const [users, setUsers] = useState([] as Array<user>)
  const toast = useToast()

  useEffect(() => {
    const f = async () => {
      const response = await fetch('/api/no_follow_users')
      const json = (await response.json())['users'] as Array<user>
      setUsers(json)
    }
    f()
  }, [])

  async function followClickHandler (event: React.MouseEvent) {
    event.preventDefault()
    const id = event.currentTarget.getAttribute('data-user-id')
    await fetch(`/api/follow/${id}`, {method: 'POST'})
    .then((response) => {
      if (response.ok) {
        toast({
          title: 'followed!',
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
                    <Button data-user-id={s.id} onClick={followClickHandler}>follow</Button>
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
