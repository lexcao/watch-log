<script>
    import EntryCard from '../component/EntryItem.svelte'
    import Socket from '../io/socket'
    import SearchBar from '../component/SearchBar.svelte'
    import Header from '../component/Header.svelte'
    import StatusBar from '../component/StatusBar.svelte'
    import { onMount } from 'svelte'

    let ws
    let connected
    let messages

    onMount(() => {
        connect()
    })

    const connect = () => {
        if ($connected) {
            return
        }
        ws = new Socket(7788)
        connected = ws.connected
        messages = ws.messages
    }
</script>

<Header/>
<StatusBar connected={$connected} onClicked={connect}/>

<SearchBar/>

{#if messages}
    {#each $messages as message}
        <EntryCard entry="{message}"/>
    {/each}
{/if}
