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

    const mock = [
        {
            Err: null,
            PipelinedObject: {
                'language': 'Go',
                'ts': 1234567891,
                'level': 'INFO',
            },
        },
        {
            Err: 'Json parse error',
            Origin: 'this is a origin log',
        },
    ]

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

{#if mock}
    {#each mock as message}
        <EntryCard entry="{message}"/>
    {/each}
{/if}
