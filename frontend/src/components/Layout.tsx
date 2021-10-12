import { Box, Flex, Text } from '@chakra-ui/react'
import React, { ReactNode } from 'react'
import { Container } from './Container'
import { Header } from './Header'
import { Sidebar } from './Sidebar'
import { Main } from './Main'

export const Layout: React.FC<{ children: ReactNode }> = ({children}) => {
  return (
    <Container>
      <Header/>

      <Flex width='100%'>
        <Box width='20%'>
          <Sidebar/>
        </Box>
        <Box width='60%'>
          <Main>
            {children}
          </Main>
        </Box>
      </Flex>

      <Flex as="footer" py="8rem">
        <Text>Footer</Text>
      </Flex>
    </Container>
  )
}
