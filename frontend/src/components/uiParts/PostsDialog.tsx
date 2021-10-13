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
import { FC, useRef } from 'react'

interface PostsDialogProps {
  handleSubmit: (text: string) => Promise<void>
  openModalButtonName: string
  headerTitle: string
}

export const PostsDialog: FC<PostsDialogProps> = (
  {handleSubmit, openModalButtonName, headerTitle}
) => {
  const {isOpen, onOpen, onClose} = useDisclosure()
  const text = useRef('')

  function originalHandleSubmit () {
    handleSubmit(text.current)
    .then(_value => {
      onClose()
    })
    .catch(reason => {
      console.error(reason)
    })
  }

  return (
    <>
      <Button onClick={onOpen}>{openModalButtonName}</Button>

      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay/>
        <ModalContent>
          <ModalHeader>{headerTitle}</ModalHeader>
          <ModalCloseButton/>
          <ModalBody pb={6}>
            <FormControl>
              <Textarea onChange={e => text.current = e.target.value}/>
            </FormControl>
          </ModalBody>

          <ModalFooter>
            <Button colorScheme='blue' mr={3} onClick={originalHandleSubmit}>
              Submit
            </Button>
            <Button onClick={onClose}>Cancel</Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  )
}
