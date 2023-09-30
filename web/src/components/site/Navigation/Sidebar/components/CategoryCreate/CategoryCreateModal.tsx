import {
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  HStack,
  Input,
  UseDisclosureProps,
  VStack,
} from "@chakra-ui/react";

import { ModalDrawer } from "src/components/site/Modaldrawer/Modaldrawer";

import { useCategoryCreate } from "./useCategoryCreate";

export function CategoryCreateModal(props: UseDisclosureProps) {
  const { register, onSubmit, errors } = useCategoryCreate(props);

  return (
    <ModalDrawer
      isOpen={props.isOpen}
      onClose={props.onClose}
      title="Create category"
    >
      <VStack as="form" onSubmit={onSubmit}>
        <FormControl>
          <FormLabel>Name</FormLabel>
          <Input {...register("name")} type="text" />
          <FormErrorMessage>{errors["name"]?.message}</FormErrorMessage>
        </FormControl>

        <FormControl>
          <FormLabel>Description</FormLabel>
          <Input {...register("description")} type="text" />
          <FormErrorMessage>{errors["description"]?.message}</FormErrorMessage>
        </FormControl>

        <HStack w="full" align="center" justify="end" pb={3} gap={4}>
          <Button variant="outline" size="sm" onClick={props.onClose}>
            Cancel
          </Button>
          <Button colorScheme="green" size="sm" type="submit">
            Create
          </Button>
        </HStack>
      </VStack>
    </ModalDrawer>
  );
}
