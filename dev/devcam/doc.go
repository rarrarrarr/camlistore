/*
Copyright 2013 The Camlistore Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
The devcam tool is a collection of wrappers around the camlistore programs
(camistored, camput, camtool...) which take care of setup and configuration,
so they can be used by developers to ease hacking on camlistore.

Usage:

	devcam <mode> [modeopts] -- [commandargs]

Modes:

  get: run camget in dev mode.
  put: run camput in dev mode.
  server: run the stand-alone camlistored in dev mode.

Examples:

  devcam get <blobref>
  devcam get -- --shared http://localhost:3169/share/<blobref>

  devcam put file --filenodes /mnt/camera/DCIM

  devcam server -wipe -mysql -fullclosure

For mode-specific help:

  devcam <mode> -help

*/
package main // import "camlistore.org/dev/devcam"
