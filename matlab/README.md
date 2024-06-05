<div align="center">
  <a href="https://matlab.mathworks.com/">
    <img alt="matlab" src="../logos/matlab.png" height="100" width="100"/>
  </a>
  <h1>Matlab</h1>
</div>


## Create launcher icon

Open a terminal and then

1. Download your own icon

```sh
sudo wget http://upload.wikimedia.org/wikipedia/commons/2/21/Matlab_Logo.png -O /usr/share/icons/matlab.png
```

2. Give your access permission

```sh
sudo touch /usr/share/applications/matlab.desktop
```

3. Edit your .desktop file

```sh
sudo gedit /usr/share/applications/matlab.desktop
```

4. Paste the following into the document

```sh
#!/usr/bin/env xdg-open
[Desktop Entry]
Type=Application
Icon=/usr/share/icons/matlab.png
Name=MATLAB R2014a
Comment=Start MATLAB - The Language of Technical Computing
Exec=matlab -desktop
Categories=Development;
```

Save, and you should have a new desktop shortcut in your launcher