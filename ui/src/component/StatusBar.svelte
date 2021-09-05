<script>
    export let connected
    export let onClicked
</script>

<div class="{connected ? 'bg-green-600' : 'bg-red-600'}">
    <div class="max-w-7xl mx-auto py-3 px-3 sm:px-6 lg:px-8">
        <div class="flex items-center justify-center">
            <p class="text-md text-white font-medium">
                {connected ? "Connected..." : "Not Connected 😭"}
            </p>
            {#if !connected}
                <div class="pl-8 mt-2 sm:order-2 sm:mt-0 sm:w-auto">
                    <button
                            on:click="{onClicked}"
                            class="px-4 py-2 rounded-md text-sm
                            text-red-400 hover:text-red-600
                            font-medium bg-white">
                        Reconnect
                    </button>
                </div>
            {/if}
        </div>
    </div>
</div>
