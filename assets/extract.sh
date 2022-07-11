rm -rf frames/
mkdir frames/
ffmpeg -i badapple.mp4 -vf fps=5/1 frames/badapple-%04d.png
