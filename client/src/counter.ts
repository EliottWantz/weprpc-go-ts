import { userService } from "./main";

export function setupCounter(element: HTMLButtonElement) {
  let counter = 0;
  const setCounter = (count: number) => {
    counter = count;
    element.innerHTML = `count is ${counter}`;
  };
  element.addEventListener("click", () => {
    userService.ping().then((res) => {
      if (res.status) {
        userService.createUser({
          username: `user${counter}`,
          password: `password${counter}`,
        });
        setCounter(counter + 1);
        userService.listUsers().then(console.log);
      }
    });
  });
  setCounter(0);
}
