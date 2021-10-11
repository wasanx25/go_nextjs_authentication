import { Box, Flex, Heading, Link as ChakraLink, List, ListItem } from '@chakra-ui/react'
import Link from 'next/link'

export const Sidebar = () => {
  return (
    <List spacing='3' padding='1rem' fontSize='1.25rem'>
      <ListItem>
        <ChakraLink>
          <Link href='/'>
            <a>Timeline</a>
          </Link>
        </ChakraLink>
      </ListItem>
      <ListItem>
        <ChakraLink>
          <Link href='/users'>
            <a>Users</a>
          </Link>
        </ChakraLink>
      </ListItem>
      <ListItem>
        <ChakraLink>
          <Link href='/follows'>
            <a>Follows</a>
          </Link>
        </ChakraLink>
      </ListItem>
    </List>
  )
}
