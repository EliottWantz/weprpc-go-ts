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

        // With generated client
        userService.listUsers().then(console.log);
        // With default fetch
        fetch("http://localhost:8081/rpc/UserService/ListUsers", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({}),
        })
          .then((res) => res.json())
          .then(console.log);
      }
    });
  });
  setCounter(0);
}
