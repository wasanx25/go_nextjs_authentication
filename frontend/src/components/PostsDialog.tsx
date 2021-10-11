import {
  Button,
  FormControl,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
  Textarea,
  useDisclosure
} from '@chakra-ui/react'
import { useState } from 'react'

export const PostsDialog = () => {
  const { isOpen, onOpen, onClose } = useDisclosure()
  const [text, setText] = useState('')

  async function handleRequest() {
    await fetch('/api/post', {
      method: 'POST',
      body: JSON.stringify({text: text})
    }).then((_response) => {
      onClose()
    }).catch((err) => {
      console.error(err)
    })
  }

  return (
    <>
      <Button onClick={onOpen}>投稿する</Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>投稿する</ModalHeader>
          <ModalCloseButton />
          <ModalBody pb={6}>
            <FormControl>
              <Textarea placeholder='投稿してみよう！' onChange={e => setText(e.target.value)}/>
            </FormControl>
          </ModalBody>

          <ModalFooter>
            <Button colorScheme="blue" mr={3} onClick={handleRequest}>
              Submit
            </Button>
            <Button onClick={onClose}>Cancel</Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  )
}
