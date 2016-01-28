/// modules - Rust
mod my {
    fn private_function() {
        println!("called `my::private_function()`");
    }

    pub fn function() {
        println!("called `my::function()`");
    }

    pub fn indirect_access() {
        print!("called `my::indirect_access()`, that\n> ");
        private_function();
    }

    // Modules can also be nested
    pub mod nested {
        pub fn function() {
            println!("called `my::nested::function()`");
        }

        #[allow(dead_code)]
        fn private_function() {
            println!("called `my::nested::private_function()`");
        }
    }
/*xcvxvsdfgv2349u23fhckd0x'sdfdsf' */
    // Nested modules follow the same rules for visibility
    mod private_nested {
        #[allow(dead_code)]
        pub fn function() {
            println!("called `my::private_nested::function()`");
        }
    }
}

fn function() {
	println!("called `function()`");
}

fn main() {
	function();
	my::function();
	
	my::indirect_access();
	my::nested::function();

  }







