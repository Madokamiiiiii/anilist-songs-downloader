# Anilist Song Downloader

This CLI script tries to download full openings and endings from anime from your anilist.
Since the approach is heuristic the results may not be perfect.

## Prerequisites
You have to install the following:
- Go (https://go.dev)
- yt-dlp (add to PATH) (https://github.com/yt-dlp/yt-dlp)
- ffmpeg (add to PATH) (https://ffmpeg.org)

## Usage
The CLI has three parameters. You can view these also with `--help`

First you have to specify your anilist username with `--username`. Make sure your profile is public.

Then you can specify with `--rescan true/false` if you want to completely renew your data 
(Default: True. If this is your first time using the script let it on).

Lastly, there's the `--getnew true/false` option. If this is true the script only gets songs which are either new or haven't been downloaded previously.
If you set this to false the script goes through every song and searches YT again. If the script gets a different search result than in a previous run it asks the user whether to use the new result or not.

## How it works
As there's no database which holds all anime theme songs we have to go the long way.
Basically the script searches YouTube for the songs.

To achieve this the scripts first loads your anilist and gets all the anime titles and MAL IDs.
Then the script uses Animethemes.org to get song information specified with the IDs.
If Animetheme doesn't find the anime the script uses MAL as a fallback.
With this information the script generates a search query and searches YT with it.
The script tries to get good results by ignoring search results with implausible length (i.e. videos longer than 10 minutes are most likely no songs).
Finally, the script downloads the result using yt-dlp and extracts the audio with ffmpeg.

The script saves the song information and also if the song has been downloaded.
Using this the script doesn't download everything again if the script is run a second time but only the things which are either new or haven't been downloaded previously.


## Problems
As both Animethemes and MAL may not have all the songs in its databases some songs may be missed.
The most complete database should be AniDB, but they don't have any API to fetch song information.
Also, there's no way to know for sure if the search result from YT actually is the right song.
In my tests most results were correct, but there were also some wrong results like videos which talk about the song or another song which happens to have the exact same title.
Lastly, the sound quality may not be perfect.
In most cases they are alright though.
