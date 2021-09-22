const router = require("express").Router();
const { default: axios } = require("axios");
const {
  models: { User, Cart, Item, CartItem },
} = require("../db");
module.exports = router;

const requireToken = async (req, res, next) => {
  try {
    req.user = await User.findByToken(req.cookies.token);
    next();
  } catch (e) {
    next(e);
  }
};

router.get("/Go", async (req, res, next) => {
  try {
    const { data } = await axios.get("http://localhost:8080/hello");
    console.log(data, "SHIT");
    res.send(data);
  } catch (err) {
    next(err);
  }
});

router.get("/", requireToken, async (req, res, next) => {
  try {
    const users = await User.findAll({
      // explicitly select only the id and username fields - even though
      // users' passwords are encrypted, it won't help if we just
      // send everything to anyone who asks!
      attributes: ["id", "username", "firstName", "lastName"],
      include: [
        {
          model: Cart,
          include: [Item],
        },
      ],
    });
    res.json(users);
  } catch (err) {
    next(err);
  }
});

router.get("/:userId", requireToken, async (req, res, next) => {
  try {
    console.log("req", req.params.userId);
    const user = await User.findOne({
      where: {
        id: req.params.userId,
      },
      include: {
        model: Cart,
        include: [Item],
      },
    });
    console.log("user", user);
    res.json(user);
  } catch (err) {
    next(err);
  }
});
