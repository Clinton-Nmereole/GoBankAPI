<section class="py-24 lg:py-28 bg-surface-900 h-full w-full overflow-hidden">
     <div class="container px-4 mx-auto h-40">
          <form class="mx-auto max-w-3xl h-40 mt-20" action="?/loginAccount" method="POST">
               <h2 class="text-6xl tracking-tight font-Satoshi py-7">Login</h2>
               
               <div class="grid grid-flow-row grid-cols-1 grid-rows-1 min-h-fit py-1">
                    <div class="md:w-4/5 w-[70%] py-5 px-1 ml-3">
                         <label for="Accountnumber">Account Number</label>
                         <input type="text" placeholder="Account Number" class="rounded-lg outline-none tracking-tight
                         text-gray-700 w-full scale-105" id="account_number" name="account_number">
                    </div>
                    <div class="md:w-4/5 w-[70%] py-5 px-1 ml-3">
                         <label for="Password">Password</label>
                         <input type="password" placeholder="Password" class="rounded-lg outline-none tracking-tight
                         text-gray-700 w-full scale-105" id="password" name="password">
                    </div>
               </div>
               <div class="w-[93%] py-8">
                    <button type="submit" class="w-[90%] bg-secondary-500 hover:bg-secondary-700 rounded-lg font-bold inline-block px-5
                    py-4"
                    >Login</button>
               </div>
               <p>Forgot Password? <a href="/forgot-password" class="text-tertiary-400 hover:underline">Reset Password</a></p>
               <p>Don't have an account? <a href="/signup" class="text-tertiary-400 hover:underline">Sign Up</a></p>
                              

          </form>

     </div>
</section>
