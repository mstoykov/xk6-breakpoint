import breakpoint from "k6/x/breakpoint"
export default function() {
    var s = "something";
    breakpoint.break({p:s})
    console.log(s)
}
