psgo.publish({ to: "a" });

let c = psgo.newSubscriber(m => {
  console.log("----------------THIS SHOULD ONLY PRINT ONCE----------------");
  let b = psgo.newSubscriber(m => {});
  psgo.subscribe(b, "b");
});
psgo.subscribe(c, "a");
