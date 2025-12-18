async function start() {
  alert(await window.go.main.App.Start(
    server.value, token.value, secret.value, listen.value
  ));
}
