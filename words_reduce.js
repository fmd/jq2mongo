use kotakuinaction;
map = function() {
    words = this.text.match(/\w+/g);
    for (index in words) {
        word = words[index].toLowerCase();
        emit(word, 1)
    }
}

reduce = function(previous, current) {
        var count = 0;

        for (index in current) {
         count += current[index];
        }

                return count;
}

result = db.runCommand({
   "mapreduce" : "comments",
       "map" : map,
       "reduce" : reduce,
       "out" : "words"})
