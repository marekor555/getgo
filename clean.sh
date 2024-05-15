echo "BUILDING PACKAGE..."
makepkg -si
echo "CLEANING AFTER INSTALL..."
rm *.tar.zst pkg src -r
echo "THANK YOU FOR USING GETGO!"
