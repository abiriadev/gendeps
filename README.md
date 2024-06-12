<div align="center">

# `gendeps`

`go mod tidy` for your JavaScript/TypeScript projects.

</div>

## Why `gendeps`?

In traditional development, we usually install dependencies first using a package manager, and then write code using these dependencies. However, the process is not always straightforward. We might install and remove multiple similar packages for experimentation, forget to remove unused dependencies after refactoring, or even forget to install some dependencies altogether.

Let's shift our perspective. Your codebase should be the single source of truth, not your `package.json` file! By embracing this concept, we can always derive `package.json` from your source code. `gendeps` makes this possible. It is a smart tool that automatically analyzes your codebase and updates the `dependencies` section in your `package.json` file.

An ideal development process with `gendeps` would look like this:

1. You write code and paste some random code from the internet.
2. `gendeps` detects new dependencies from the added code and updates your `package.json` file accordingly.
3. Optionally, `gendeps` can run your favorite package manager to install the new dependencies automatically.
4. You then refactor your code, potentially removing some dependencies.
5. `gendeps` detects unused dependencies and removes them from your `package.json` file.
6. Optionally, `gendeps` can run your favorite package manager to remove the unused dependencies automatically.
7. When you're ready to commit your changes, your `package.json` will be clean and only contain the dependencies you are using. Simply run `git add package.json`, and you're done!
