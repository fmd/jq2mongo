use kotakuinaction;
db.posts.ensureIndex({title: "text"});
db.comments.ensureIndex({text: "text"});
