/** @type {import('tailwindcss').Config} */
module.exports = {
	content: [
		"./pages/**/*.{js,ts,jsx,tsx}",
		"./components/**/*.{js,ts,jsx,tsx}",
		"./app/**/*.{js,ts,jsx,tsx}",
	],
	corePlugins: {
		aspectRatio: false, // avoid conflicts with native tailwind util
	},

	plugins: [require("@tailwindcss/forms")],
	plugins: [require("@tailwindcss/aspect-ratio")],
};
