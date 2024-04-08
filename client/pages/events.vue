<template>
  <main id="site__main" class="2xl:ml-[--w-side] xl:ml-[--w-side-sm] p-2.5 h-[calc(100vh-var(--m-top))] mt-[--m-top]"
    style="max-height: 100%; overflow-y: auto;">
    <div class="2xl:max-w-[1220px] max-w-[1065px] mx-auto">
      <div class="page-heading">
        <h1 class="page-title">
          Events
        </h1>

        <nav class="nav__underline">
          <ul uk-tab class="group"
            uk-switcher="connect: #ttabs ; animation: uk-animation-slide-right-medium, uk-animation-slide-left-medium">
            <li> <a href="#"> Suggestions </a> </li>
            <li> <a href="#"> Popular </a> </li>
            <li> <a href="#"> My events </a> </li>
          </ul>
        </nav>
      </div>

      <!-- event feautred -->
      <div class="relative" tabindex="-1" uk-slider="finite:true">
        <div class="uk-slider-container pb-1">
          <ul class="uk-slider-items grid-small">
            <li class="lg:w-1/4 sm:w-1/3 w-1/2">
              <div class="card">
                <a href="timeline-event.html">
                  <div class="card-media h-32">
                    <img src="../assets/images/post/img-3.jpg" alt="">
                    <div class="card-overly" />
                  </div>
                </a>
                <div class="card-body">
                  <p class="text-xs font-medium text-blue-600 mb-1">
                    Next week
                  </p>
                  <a href="timeline-event.html">
                    <h4 class="card-title text-sm"> About Safety and Flight </h4>
                  </a>
                  <a href="#">
                    <p class="card-text text-black mt-2"> Dubai </p>
                  </a>
                  <div class="card-list-info text-xs mt-1">
                    <div> 26 Intersted</div>
                    <div class="md:block hidden">
                      ·
                    </div>
                    <div> 8 Going</div>
                  </div>
                  <div class="flex gap-2">
                    <button type="button" class="button bg-primary text-white flex-1">
                      Intersted
                    </button>
                    <button type="button" class="button bg-secondery !w-auto">
                      <i class='bx bx-redo text-lg' ></i>
                    </button>
                  </div>
                </div>
              </div>
            </li>
            <li class="lg:w-1/4 sm:w-1/3 w-1/2">
              <div class="card">
                <a href="timeline-event.html">
                  <div class="card-media h-32">
                    <img src="../assets/images/post/img-2.jpg" alt="">
                    <div class="card-overly" />
                  </div>
                </a>
                <div class="card-body">
                  <p class="text-xs font-semibold text-teal-600 mb-1">
                    Opening
                  </p>
                  <a href="timeline-event.html">
                    <h4 class="card-title text-sm"> Wedding trend Ideas </h4>
                  </a>
                  <a href="#">
                    <p class="card-text text-black mt-2"> Turkey </p>
                  </a>
                  <div class="card-list-info text-xs mt-1">
                    <div> 20 Intersted</div>
                    <div class="md:block hidden">
                      ·
                    </div>
                    <div> 16 Going</div>
                  </div>
                  <div class="flex gap-2">
                    <button type="button" class="button bg-primary text-white flex-1">
                      Intersted
                    </button>
                    <button type="button" class="button bg-secondery !w-auto">
                      <i class='bx bx-redo text-lg' ></i>
                    </button>
                  </div>
                </div>
              </div>
            </li>
            <li class="lg:w-1/4 sm:w-1/3 w-1/2">
              <div class="card">
                <a href="timeline-event.html">
                  <div class="card-media h-32">
                    <img src="../assets/images/post/post-1.jpg" alt="">
                    <div class="card-overly" />
                  </div>
                </a>
                <div class="card-body">
                  <p class="text-xs font-medium text-red-600 mb-1">
                    WED JUL 10,2024 AT 10PM
                  </p>
                  <a href="timeline-event.html">
                    <h4 class="card-title text-sm"> The global creative </h4>
                  </a>
                  <a href="#">
                    <p class="card-text text-black mt-2"> Japan </p>
                  </a>
                  <div class="card-list-info text-xs mt-1">
                    <div> 15 Intersted</div>
                    <div class="md:block hidden">
                      ·
                    </div>
                    <div> 2 Going</div>
                  </div>
                  <div class="flex gap-2">
                    <button type="button" class="button bg-primary text-white flex-1">
                      Intersted
                    </button>
                    <button type="button" class="button bg-secondery !w-auto">
                      <i class='bx bx-redo text-lg' ></i>
                    </button>
                  </div>
                </div>
              </div>
            </li>
            <li class="lg:w-1/4 sm:w-1/3 w-1/2">
              <div class="card">
                <a href="timeline-event.html">
                  <div class="card-media h-32">
                    <img src="../assets/images/post/img-4.jpg" alt="">
                    <div class="card-overly" />
                  </div>
                </a>
                <div class="card-body">
                  <p class="text-xs font-semibold text-teal-600 mb-1">
                    Opening
                  </p>
                  <a href="timeline-event.html">
                    <h4 class="card-title text-sm"> Perspective is everything </h4>
                  </a>
                  <a href="#">
                    <p class="card-text text-black mt-2"> London </p>
                  </a>
                  <div class="card-list-info text-xs mt-1">
                    <div> 20 Intersted</div>
                    <div class="md:block hidden">
                      ·
                    </div>
                    <div> 16 Going</div>
                  </div>
                  <div class="flex gap-2">
                    <button type="button" class="button bg-primary text-white flex-1">
                      Intersted
                    </button>
                    <button type="button" class="button bg-secondery !w-auto">
                      <i class='bx bx-redo text-lg' ></i>
                    </button>
                  </div>
                </div>
              </div>
            </li>
            <li class="lg:w-1/4 sm:w-1/3 w-1/2">
              <div class="card">
                <a href="timeline-event.html">
                  <div class="card-media h-32">
                    <img src="../assets/images/post/img-3.jpg" alt="">
                    <div class="card-overly" />
                  </div>
                </a>
                <div class="card-body">
                  <p class="text-xs font-medium text-blue-600 mb-1">
                    Next week
                  </p>
                  <a href="timeline-event.html">
                    <h4 class="card-title text-sm"> About Safety and Flight </h4>
                  </a>
                  <a href="#">
                    <p class="card-text text-black mt-2"> Dubai </p>
                  </a>
                  <div class="card-list-info text-xs mt-1">
                    <div> 26 Intersted</div>
                    <div class="md:block hidden">
                      ·
                    </div>
                    <div> 8 Going</div>
                  </div>
                  <div class="flex gap-2">
                    <button type="button" class="button bg-primary text-white flex-1">
                      Intersted
                    </button>
                    <button type="button" class="button bg-secondery !w-auto">
                      <i class='bx bx-redo text-lg' ></i>
                    </button>
                  </div>
                </div>
              </div>
            </li>
          </ul>
        </div>

        <!-- slide nav icons -->
        <a class="nav-prev !top-20" href="#" uk-slider-item="previous"> <i class='bx bx-chevron-left text-2xl' ></i> </a>
        <a class="nav-next !top-20" href="#" uk-slider-item="next"> <i class='bx bx-chevron-right text-2xl' ></i> </a>
      </div>

      <div class="sm:my-6 my-3 flex items-center justify-between md:mt-10">
        <div>
          <h2 class="text-xl font-semibold">
            Lists You May Like
          </h2>
          <p class="font-normal text-sm text-gray-500 leading-6">
            Find a group by browsing top categories.
          </p>
        </div>
        <a href="#" class="text-blue-500 sm:block hidden text-sm"> See all </a>
      </div>

      <!-- listing  slider -->
      <div class="grid lg:grid-cols-1 md:grid-cols-5 sm:grid-cols-4 gap-2.5 mt-8">
        <div class="uk-slider-container pb-1">
          <ul class="uk-slider-items grid-small">
            <li class="md:w-1/5 sm:w-1/3 w-1/2">
              <a href="#">
                <div class="relative rounded-lg overflow-hidden">
                  <img src="../assets/images/post/img-2.jpg" alt="" class="h-36 w-full object-cover">
                  <div class="w-full bottom-0 absolute left-0 bg-gradient-to-t from-black/60 pt-10">
                    <div class="text-white p-5">
                      <div class="text-sm font-light"> Miami </div>
                      <div class="text-lg leading-3 mt-1.5"> Hotels </div>
                    </div>
                  </div>
                </div>
              </a>
            </li>
            <li class="md:w-1/5 sm:w-1/3 w-1/2">
              <a href="#">
                <div class="relative rounded-lg overflow-hidden">
                  <img src="../assets/images/post/img-3.jpg" alt="" class="h-36 w-full object-cover">
                  <div class="w-full bottom-0 absolute left-0 bg-gradient-to-t from-black/60 pt-10">
                    <div class="text-white p-5">
                      <div class="text-sm font-light"> Florida </div>
                      <div class="text-lg leading-3 mt-1.5"> Hotels </div>
                    </div>
                  </div>
                </div>
              </a>
            </li>
            <li class="md:w-1/5 sm:w-1/3 w-1/2">
              <a href="#">
                <div class="relative rounded-lg overflow-hidden">
                  <img src="../assets/images/post/img-4.jpg" alt="" class="h-36 w-full object-cover">
                  <div class="w-full bottom-0 absolute left-0 bg-gradient-to-t from-black/60 pt-10">
                    <div class="text-white p-5">
                      <div class="text-sm font-light"> London </div>
                      <div class="text-lg leading-3 mt-1.5"> Hotels </div>
                    </div>
                  </div>
                </div>
              </a>
            </li>
            <li class="md:w-1/5 sm:w-1/3 w-1/2">
              <a href="#">
                <div class="relative rounded-lg overflow-hidden">
                  <img src="../assets/images/post/post-1.jpg" alt="" class="h-36 w-full object-cover">
                  <div class="w-full bottom-0 absolute left-0 bg-gradient-to-t from-black/60 pt-10">
                    <div class="text-white p-5">
                      <div class="text-sm font-light"> Dubai </div>
                      <div class="text-lg leading-3 mt-1.5"> Hotels </div>
                    </div>
                  </div>
                </div>
              </a>
            </li>
            <li class="md:w-1/5 sm:w-1/3 w-1/2">
              <a href="#">
                <div class="relative rounded-lg overflow-hidden">
                  <img src="../assets/images/post/img-2.jpg" alt="" class="h-36 w-full object-cover">
                  <div class="w-full bottom-0 absolute left-0 bg-gradient-to-t from-black/60 pt-10">
                    <div class="text-white p-5">
                      <div class="text-sm font-light"> Turkey </div>
                      <div class="text-lg leading-3 mt-1.5"> Resturent </div>
                    </div>
                  </div>
                </div>
              </a>
            </li>
            <li class="md:w-1/5 sm:w-1/3 w-1/2">
              <a href="#">
                <div class="relative rounded-lg overflow-hidden">
                  <img src="../assets/images/post/img-2.jpg" alt="" class="h-36 w-full object-cover">
                  <div class="w-full bottom-0 absolute left-0 bg-gradient-to-t from-black/60 pt-10">
                    <div class="text-white p-5">
                      <div class="text-sm font-light"> Miami </div>
                      <div class="text-lg leading-3 mt-1.5"> Hotels </div>
                    </div>
                  </div>
                </div>
              </a>
            </li>
          </ul>
        </div>

        <!-- slide nav icons -->
        <a class="nav-prev" href="#" uk-slider-item="previous"> <i class='bx bx-chevron-left text-2xl' ></i> </a>
        <a class="nav-next" href="#" uk-slider-item="next"> <i class='bx bx-chevron-right text-2xl' ></i> </a>
      </div>

      <div class="flex items-center justify-between text-black dark:text-white py-3 mt-6">
        <h3 class="text-xl font-semibold">
          Upcomming Events
        </h3>
        <a href="#" class="text-sm text-blue-500">See all</a>
      </div>

      <!-- event grid -->
      <div class="grid lg:grid-cols-4 md:grid-cols-3 sm:grid-cols-2 gap-2.5 mt-4">
        <div class="card">
          <a href="timeline-event.html">
            <div class="card-media h-32">
              <img src="../assets/images/post/post-1.jpg" alt="">
              <div class="card-overly" />
            </div>
          </a>
          <div class="card-body">
            <p class="text-xs font-medium text-black text-red-600 mb-1">
              WED JUL 10,2024 AT 10PM
            </p>
            <a href="timeline-event.html">
              <h4 class="card-title text-sm"> The global creative </h4>
            </a>
            <a href="#">
              <p class="card-text text-black mt-2"> Japan </p>
            </a>
            <div class="card-list-info text-xs mt-1">
              <div> 15 Intersted</div>
              <div class="md:block hidden">
                ·
              </div>
              <div> 2 Going</div>
            </div>
            <div class="flex gap-2">
              <button type="button" class="button bg-primary text-white flex-1">
                Intersted
              </button>
              <button type="button" class="button bg-secondery !w-auto">
                <i class='bx bx-redo text-lg' ></i>
              </button>
            </div>
          </div>
        </div>

        <div class="card">
          <a href="timeline-event.html">
            <div class="card-media h-32">
              <img src="../assets/images/post/img-2.jpg" alt="">
              <div class="card-overly" />
            </div>
          </a>
          <div class="card-body">
            <p class="text-xs font-semibold text-teal-600 mb-1">
              Opening
            </p>
            <a href="timeline-event.html">
              <h4 class="card-title text-sm"> Wedding trend Ideas </h4>
            </a>
            <a href="#">
              <p class="card-text text-black mt-2"> Turkey </p>
            </a>
            <div class="card-list-info text-xs mt-1">
              <div> 20 Intersted</div>
              <div class="md:block hidden">
                ·
              </div>
              <div> 16 Going</div>
            </div>
            <div class="flex gap-2">
              <button type="button" class="button bg-primary text-white flex-1">
                Intersted
              </button>
              <button type="button" class="button bg-secondery !w-auto">
                <i class='bx bx-redo text-lg' ></i>
              </button>
            </div>
          </div>
        </div>

        <div class="card">
          <a href="timeline-event.html">
            <div class="card-media h-32">
              <img src="../assets/images/post/img-3.jpg" alt="">
              <div class="card-overly" />
            </div>
          </a>
          <div class="card-body">
            <p class="text-xs font-medium text-black text-red-600 mb-1">
              WED JUL 10,2024 AT 10PM
            </p>
            <a href="timeline-event.html">
              <h4 class="card-title text-sm"> About Safety and Flight </h4>
            </a>
            <a href="#">
              <p class="card-text text-black mt-2"> Dubai </p>
            </a>
            <div class="card-list-info text-xs mt-1">
              <div> 26 Intersted</div>
              <div class="md:block hidden">
                ·
              </div>
              <div> 8 Going</div>
            </div>
            <div class="flex gap-2">
              <button type="button" class="button bg-primary text-white flex-1">
                Intersted
              </button>
              <button type="button" class="button bg-secondery !w-auto">
                <i class='bx bx-redo text-lg' ></i>
              </button>
            </div>
          </div>
        </div>

        <div class="card">
          <a href="timeline-event.html">
            <div class="card-media h-32">
              <img src="../assets/images/post/img-4.jpg" alt="">
              <div class="card-overly" />
            </div>
          </a>
          <div class="card-body">
            <p class="text-xs font-semibold text-teal-600 mb-1">
              Opening
            </p>
            <a href="timeline-event.html">
              <h4 class="card-title text-sm"> Perspective is everything </h4>
            </a>
            <a href="#">
              <p class="card-text text-black mt-2"> London </p>
            </a>
            <div class="card-list-info text-xs mt-1">
              <div> 20 Intersted</div>
              <div class="md:block hidden">
                ·
              </div>
              <div> 16 Going</div>
            </div>
            <div class="flex gap-2">
              <button type="button" class="button bg-primary text-white flex-1">
                Intersted
              </button>
              <button type="button" class="button bg-secondery !w-auto">
                <i class='bx bx-redo text-lg' ></i>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </main>
</template>

<script setup>
definePageMeta({
  middleware: ["auth-only"],
})
</script>