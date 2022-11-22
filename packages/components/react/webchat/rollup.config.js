import resolve from "@rollup/plugin-node-resolve"
import commonjs from "@rollup/plugin-commonjs"
import typescript from "@rollup/plugin-typescript"
import url from "@rollup/plugin-url"
import dts from "rollup-plugin-dts"
import postcss from 'rollup-plugin-postcss'
import postcssImport from 'postcss-import'
import autoprefixer from 'autoprefixer'

const packageJson = require("./package.json")

export default [
    {
        input: "src/index.ts",
        output: [
            {
                file: packageJson.main,
                format: "cjs",
                sourcemap: true,
            },
            {
                file: packageJson.module,
                format: "esm",
                sourcemap: true,
            },
        ],
        plugins: [
            resolve(),
            url(),
            commonjs(),
            typescript({tsconfig: "./tsconfig.json"}),
            postcss({
                extract: true,
                plugins: [postcssImport(), autoprefixer()],
            }),
        ],
        external: ["react", "react-dom", "styled-components"]
    },
    {
        input: "dist/esm/index.d.ts",
        output: [{file: "dist/index.d.ts", format: "esm"}],
        plugins: [dts()],
        external: [/\.css$/],
    },
]