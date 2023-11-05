# Exterminator
CLI tool that wraps and empowers **CursedSpirits** with some hot-reloads and amplifications.
> Keep in mind; Exterminator is for educational purposes only and I am not responsable for anything you or it may cause.

### Features
- [ ] Installer to install CursedSpirits globally
- [x] Refresh CursedSpirits ongoing attacks automatically with new proxies and the given interval refresh ratio; allowing you to down a website limitless, even if the proxies get's banned

<img src="https://i0.wp.com/anitrendz.net/news/wp-content/uploads/2023/08/ragnacrimson_trailer3screenshot.png?fit=1920%2C1080&ssl=1" width="600">

> **Yet the power is again derived from an anime. Empower CursedSpirits with some silverine. Mix up Gojo with his teenage mutant version^^.**

# Supported OS
Works on Debain based linux distributions

# Installation
> Install CursedSpirits into ``$HOME/CursedSpirits`` and navigate to this directory then
> build CursedSpirits with ``go build``
>
> Now install Exterminator into ``$HOME/exterminator`` and navigate to this directory then build Exterminator with ``go build``
> 
> Move the binary program from the go build output in the exterminator folder to CursedSpirits folder:<br>
> ``mv $HOME/exterminator $HOME/CursedSpirits``
>
> **All done, read the usage instructions**

# Usage
Keep in mind that the hot-reload uses public proxies and not local ones, soon a pipe stream functionality for custom proxies will be released.

``./exterminator --cmd "--url https://target.com" --refresh 6 --bin ./CursedSpirits``
