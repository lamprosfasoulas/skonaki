# Skonaki

Skonaki is a cheatsheet api written in go. It uses community based libraries that you can see bellow.

## How to use

Skonaki can be used either with curl or with a browser.
Type `curl example.com/{ query }` and you will see the results.
For the brower version you can either type the URL or use the integrated form that uses HTMX.

Type `curl example.com/{ query }/{ specific_topic }` for more specific things.

Type `curl example.com/:list` to get the list of all the available cheat sheets.

## How to Host

Skonaki can be hosted with Docker Compose or you can compile the source code yourself.

## How it works 

All the cheat sheets are stored under the `data/` folder. The numbers in front of the folders is used 
for the display order.

The app uses some env variables. Those are:

| ENV VAR | Value| USE |
| :---: | :---: |:---:|
|SKON_REDIS_ADDR |**Default:** `localhost:6379` | Redis Cache Server Address|
|SKON_REDIS_PASSWD|**Default:** `empty` | Redis Cache Server Password|
|SKON_ALLOW_API|`true` or `false` or `empty` | Enables the file update API|
|SKON_ALLOW_SUGGEST|`true` or `false` or `empty`|Enables the suggestions form|
|SKON_DOMAIN |**Default**: localhost:42069  | The domain that will be displayed to the user|

****empty is the same as false***

## Redis

The app uses a simple redis instance to store the results as a byte array. The cached data
have a KeepTTL of 5 hours. Up until now there is not cache invalidation. It just expires after 5 hours.

## Suggestion form

By setting the env variable to **true** (this is text) you enable the `example.com/:suggest` endpoint where you can make new suggestions.
The page has two inputs, one is for the query and one if for the result body. 

If for example you want to suggest `example.com/thislink` then you would type `thislink` in the input and whatever you want in the body. Keep in mind that 
the text will get syntax highlighting. ( The default language is bash ). So mind the syntax of you suggestion.

If you want to suggest `example.com/thislink/sublink` you have to type ` thislink/sublink` in the form input.

The app looks for the requested query in every folder under `data/`, if it finds a matching file, it returns it otherwise it ignores it. For 
queries like this `example.com/thislink/sublink` the app look for `_thislink/sublink` in each folder inside `data/`

The suggested files are saved in the `suggestions` directory.

## API 

By setting the env variable to **true** (yes text) you enable the `example.com/:api` endpoint. There you can perform POST requests. 

The request form must include a **path** and **content** field.

When an api call is performed the app creates a files inside the folder `11.internal`, uses the form
**path** as path, and writes the **content** inside the file. If you give the name of a file that exists, the file will be overwritten.

## Config

There is a config folder inside `data/` there you can put files like 404 error and home.

## Sidenote 

This app uses template to render both the terminal and the html responses so the domain for example
is set this way.

---

## Project sources
| Repository | Link |
| ----------  | ---- |
| tldr-pages/tldr| [Repo](https://github.com/tldr-pages/tldr) |
| cheat/cheatsheets | [Repo](https://github.com/cheat/cheatsheets)
| chubin/cheat.sheets | [Repo](https://github.com/chubin/cheat.sheets) |

---
> Some components of this project are based on public domain works licensed under CC0 1.0 Universal.

> This project includes work licensed under [Creative Commons Attribution 4.0 International License](https://creativecommons.org/licenses/by/4.0/). 
> Original work available at [Repo](https://github.com/tldr-pages/tldr). 



> This project includes software licensed under the MIT License by [Igor Chubin].
> Original license:
>
> MIT License
>
> Copyright (c) 2017 Igor Chubin
>
> Permission is hereby granted, free of charge, to any person obtaining a copy
> of this software and associated documentation files (the "Software"), to deal
> in the Software without restriction, including without limitation the rights
> to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
> copies of the Software, and to permit persons to whom the Software is
> furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included in all
> copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
> SOFTWARE.

