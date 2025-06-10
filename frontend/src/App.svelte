<script>
  import { Login, DeleteUser, RegisterUser } from '../wailsjs/go/main/App.js'
  import { Quit } from '../wailsjs/runtime/runtime.js'

  let username = ""
  let password = ""
  let userData = { name: "", permission: "" }
  let loginError = ""
  let deleteTarget = ""
  let deleteMessage = ""
  let currentPage = "login"

  async function attemptLogin() {
    const result = await Login(username, password)
    if (result.success) {
      userData = result
      currentPage = "home"
      loginError = ""
    } else {
      loginError = "❌ Nome utente o password errati."
    }
  }

  async function deleteUser() {
    const ok = await DeleteUser(deleteTarget)
    deleteMessage = ok
      ? `✅ Utente '${deleteTarget}' eliminato.`
      : `❌ Impossibile eliminare '${deleteTarget}'.`
    deleteTarget = ""
  }

  async function registerUser() {
    const ok = await RegisterUser(username, password)
    if (ok) {
      alert("✅ Registrazione completata! Ora puoi fare login.")
      username = ""
      password = ""
      currentPage = "login"
    } else {
      alert("❌ Registrazione fallita: utente già esistente o errore.")
    }
  }

  function logout() {
    username = ""
    password = ""
    userData = { name: "", permission: "" }
    currentPage = "login"
    loginError = ""
    deleteMessage = ""
  }

  function goToRegister() {
    currentPage = "register"
  }

  function backToLogin() {
    currentPage = "login"
  }

  function closeApp() {
    Quit()
  }
</script>

{#if currentPage === 'login'}
  <main>
    <h2 class="text">Login</h2>
    <input bind:value={username} placeholder="Username" />
    <input bind:value={password} type="password" placeholder="Password" />
    <button on:click={attemptLogin}>Login</button>
    {#if loginError}
      <div>{loginError}</div>
    {/if}
    <p class="text">
      Non hai un account? <a href="#" on:click|preventDefault={goToRegister}>Registrati</a>
    </p>
  </main>

{:else if currentPage === 'register'}
  <main>
    <h2 class="text">Registrazione</h2>
    <input bind:value={username} placeholder="Nuovo username" />
    <input bind:value={password} type="password" placeholder="Password" />
    <button on:click={registerUser}>Crea Account</button>
    <p class="text">
      Hai già un account? <a href="#" on:click|preventDefault={backToLogin}>Torna al login</a>
    </p>
  </main>

{:else if currentPage === 'home'}
  <main>
    <h2 class="text">Benvenuto, {userData.name}!</h2>
    <p class="text">Permesso: <strong>{userData.permission}</strong></p>
    <button on:click={logout}>Logout</button>

    {#if userData.permission === "admin"}
      <section>
        <h3 class="text">Gestione Utenti (Admin)</h3>
        <input bind:value={deleteTarget} placeholder="Utente da eliminare" />
        <button on:click={deleteUser}>Elimina Utente</button>
        {#if deleteMessage}
          <p class="text">{deleteMessage}</p>
        {/if}
      </section>
    {:else}
      <p class="text">Accesso base. Nessun controllo amministrativo disponibile.</p>
    {/if}
  </main>
{/if}

<button class="close-btn" on:click={closeApp}>Chiudi App</button>