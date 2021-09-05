<script>
    export let entry
    const { PipelinedObject, Origin } = entry
</script>

<div class="shadow m-auto hover:shadow-md rounded w-10/12 my-4 p-4">
    {#if !entry.Err}
        {#each Object.keys(PipelinedObject) as key}
            <div class="flex divide-x-4 hover:shadow-lg items-center">
                <div class="px-4 py-2 text-sm font-bold text-gray-600">
                    {key}
                </div>
                <div class="flex-grow px-4 bg-white">
                    {PipelinedObject[key]}
                    123213123213123123123123123
                    123123
                    123
                    12312312
                </div>
            </div>
        {/each}
    {:else}
        <div class="bg-red-200">{entry.Err} : {Origin}</div>
    {/if}
</div>
