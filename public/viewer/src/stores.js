import { writable } from "svelte/store";

export const visible = writable(true);

export const currentRoom = writable({
  id: 1,
  title: "test title",
  description: `description description`,
  actions: []
});

export function setCurrentRoom(id) {
  const room = rooms.find(r => r.id === id);
  currentRoom.update(r => room);
}

export const rooms = [
  {
    id: 1,
    title: "a new beginning",
    description: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris iaculis, dui at lacinia laoreet, eros ligula elementum ligula, vitae interdum tellus urna sed ipsum. Nullam luctus, massa in tristique posuere, augue enim blandit tellus, in dictum mi est nec arcu. Aliquam aliquam mi in odio sollicitudin venenatis. Suspendisse potenti. Nunc sem mi, sagittis nec vulputate a, semper vel augue. Maecenas laoreet sem vitae urna tempus, eu ullamcorper leo euismod. Morbi vulputate, enim eget hendrerit tristique, elit odio facilisis leo, sed pulvinar metus sapien at ligula. Pellentesque vehicula blandit sollicitudin.
	Lorem ipsum dolor sit amet, quis luctus est hendrerit in. Vestibulum non malesuada lorem, a porttitor nulla. Morbi molestie euismod cursus. Nunc at tellus id nunc pharetra pretium. Aliquam gravida, tortor sed commodo lobortis, nisi neque posuere lectus, fringilla congue orci tellus laoreet sapien.`,
    actions: [
      {
        type: "direction",
        description: "Go North",
        data: {
          room: 2
        }
      },
      {
        type: "direction",
        description: "Go through east exit",
        data: {
          room: 2
        }
      }
    ]
  },
  {
    id: 2,
    title: "A secret passage",
    description: `Loremit amet, quis luctus est hendrerit in. Vestibulum non malesuada lorem, a porttitor nulla. Morbi molestie euismod cursus. Nunc at tellus id nunc pharetra pretium. Aliquam gravida, tortor sed commodo lobortis, nisi neque posuere lectus, fringilla congue orci tellus laoreet sapien.`,
    actions: [
      {
        type: "direction",
        description: "Go Back",
        data: {
          room: 1
        }
      }
    ]
  }
];
