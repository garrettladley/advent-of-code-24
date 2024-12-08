pub mod p1;
// See https://github.com/rust-lang/rust-clippy/issues/13185
// Using map on L25 is intentional as we need to mutate the Vec<u32> in place
#[allow(clippy::manual_inspect)]
pub mod p2;
