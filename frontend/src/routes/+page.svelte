<script lang="ts">
  import { fetchAppointments } from '$lib/api.service';
  import { onMount } from 'svelte';

  let appointments = $state<{ State: string }[]>([]);

  onMount(async () => {
    const response = await fetchAppointments();
    console.log('on component', response);

    if (response.failure || !response.data.appointments) {
      return;
    }

    appointments = response.data.appointments;
  });
</script>

<h1>Welcome to SvelteKit</h1>
<p>Visit <a href="https://svelte.dev/docs/kit">svelte.dev/docs/kit</a> to read the documentation</p>

{#each appointments as appointment}
  <h2>{appointment.State}</h2>
{/each}
