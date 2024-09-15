/** @type {import('tailwindcss').Config} */
export default {
  content: ["./routes/**/*.templ"],
  theme: {
    extend: {
      gridTemplateRows: {
        layout: "1fr 50px",
      },
    },
  },
  daisyui: {
    themes: ["light", "dark", "cupcake", "cyberpunk"],
  },
  plugins: [require("daisyui")],
};
