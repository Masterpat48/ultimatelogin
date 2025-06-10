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
      loginError = "❌ Username or Password are wrong."
    }
  }

  async function deleteUser() {
    const ok = await DeleteUser(deleteTarget)
    deleteMessage = ok
      ? `✅ User '${deleteTarget}' deleted.`
      : `❌ Can't delete user '${deleteTarget}'.`
    deleteTarget = ""
  }

  async function registerUser() {
    const ok = await RegisterUser(username, password)
    if (ok) {
      alert("✅ Signin compleated, now you can login.")
      username = ""
      password = ""
      currentPage = "login"
    } else {
      alert("❌ Signin failed, Username already exist or database error.")
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
      <div class="text">{loginError}</div>
    {/if}
    <p class="text">
      Don't have an account? <a href="#" on:click|preventDefault={goToRegister}>Singin</a>
    </p>
  </main>

{:else if currentPage === 'register'}
  <main>
    <h2 class="text">Signin</h2>
    <input bind:value={username} placeholder="Username" />
    <input bind:value={password} type="password" placeholder="Password" />
    <button on:click={registerUser}>Create Account</button>
    <p class="text">
      Already signed in? <a href="#" on:click|preventDefault={backToLogin}>Back to the login</a>
    </p>
  </main>

{:else if currentPage === 'home'}
  <main>
    <h2 class="text">Welcome, {userData.name}!</h2>
    <p class="text">Permission: <strong>{userData.permission}</strong></p>
    <button on:click={logout}>Logout</button>

    {#if userData.permission === "admin"}
      <section>
        <h3 class="text">Users administration (Admin)</h3>
        <input bind:value={deleteTarget} placeholder="User to delete" />
        <button on:click={deleteUser}>Delete User</button>
        {#if deleteMessage}
          <p class="text">{deleteMessage}</p>
        {/if}
      </section>
    {:else}
      <p class="text">Base access, no administration permits.</p>
    {/if}
  </main>
{/if}

<button class="close-btn" on:click={closeApp}>Quit app</button>