for file in *.png; do
    echo "$file"
	convert "$file" -trim +repage "$file" 
	convert "$file"           \( +clone -rotate 90 +clone -mosaic +level-colors "rgb(184,204,8)" \)           +swap -gravity center -background none -composite    "$file"
	convert "$file" -transparent "rgb(184,204,8)" "$file"
done

echo "Adding standard padding..."
mogrify -path "./" -bordercolor transparent -border 100 -format png *.png
echo "Done!"
