module.exports = {
    mode: process.env.NODE_ENV && 'jit',
    purge: [
        './public/index.html',
        './src/**/*.svelte',
    ],
    darkMode: false, // or 'media' or 'class'
    theme: {
        extend: {},
    },
    variants: {},
    plugins: [],
}
