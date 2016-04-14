fn main()
{
  let wflg = 0;
  let tflg = 0;
  let dflg = 0;
  let c = 6;
  match c {
    3=>wflg = 1   ,
      4=>              wflg = 1 ,
        5=>    tflg = 1,
         6=>    tflg = 45,
            7=> dflg = 1,
    }
  printf("match %d",tflg);
    return 0;
}